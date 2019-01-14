package notify;
 
import (
    "github.com/fsnotify/fsnotify"
    "github.com/smtp-http/filenotify_go/conn"
    "log"
    "fmt"
    "path/filepath"
    "path"
    "io/ioutil"
)

type FileMonitor struct {
    m_tcpserver *conn.TcpServer
}
 
func (f *FileMonitor)Monitor() {



    //创建一个监控对象
    watch, err := fsnotify.NewWatcher();
    if err != nil {
        log.Fatal(err);
    }
    defer watch.Close();
    //添加要监控的对象，文件或文件夹
    err = watch.Add("./tmp");
    if err != nil {
        log.Fatal(err);
    }
    //我们另启一个goroutine来处理监控对象的事件
    go func() {
        for {
            select {
            case ev := <-watch.Events:
                {
                    //判断事件发生的类型，如下5种
                    // Create 创建
                    // Write 写入
                    // Remove 删除
                    // Rename 重命名
                    // Chmod 修改权限
                    if ev.Op&fsnotify.Create == fsnotify.Create {
                        fmt.Println("创建文件 : ", ev.Name);
                    }
                    if ev.Op&fsnotify.Write == fsnotify.Write {
                        fmt.Println("写入文件 : ", ev.Name);
                        paths, fileName := filepath.Split(ev.Name) 
                        fmt.Println(paths, fileName) //获取路径中的目录及文件名 E:\data\ test.txt 
                        fmt.Println(filepath.Base(ev.Name)) //获取路径中的文件名test.txt 
                        

                        if path.Ext(ev.Name) == ".txt" {
                            fmt.Println(path.Ext(ev.Name)) //获取路径中的文件的后缀 .txt
                            b, err := ioutil.ReadFile(ev.Name) 
                            if err != nil { 
                                fmt.Print(err) 
                            } 
                            fmt.Println(b) 
                            //str := string(b) 
                        
                            f.m_tcpserver.Notify(b)
                        }

                        
                    }

                    if ev.Op&fsnotify.Remove == fsnotify.Remove {
                        fmt.Println("删除文件 : ", ev.Name);
                    }

                    if ev.Op&fsnotify.Rename == fsnotify.Rename {
                        fmt.Println("重命名文件 : ", ev.Name);
                    }
                    
                    if ev.Op&fsnotify.Chmod == fsnotify.Chmod {
                        fmt.Println("修改权限 : ", ev.Name);
                    }
                }
            case err := <-watch.Errors:
                {
                    log.Println("error : ", err);
                    return;
                }
            }
        }
    }();
 
    //循环
    select {};
}
