package main

import (
    "fmt"
    "net/http"
    "log"
    "io/ioutil"
    //"encoding/json"
    "strconv"
    "math"
    "runtime"
    "github.com/pquerna/ffjson/ffjson"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	resp,err := http.Get("http://localhost:18082/product_list.json")
    if err!=nil{
        fmt.Println(err.Error())
        return
    }
    body, _ := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()

    //bodyStr := string(body)
    //json_bytes := []byte(bodyStr)

    var m interface{}
    ffjson.Unmarshal(body,&m)
    data := m.(map[string]interface{})

    for _,v := range data["list"].([]interface{}) {

        d := v.(map[string]interface{})
        value,_ := strconv.ParseFloat( d["price"].(map[string]interface{})["value"].(string),64 )
        d["price"].(map[string]interface{})["value"] = strconv.FormatFloat((value*1.25),'f',-1,64)
        
        var color_quantity int64
        for _,o_interf := range d["options"].([]interface{}) {
            o := o_interf.(map[string]interface{})
            if(o["title"] == "Color"){
                color_quantity,_ = strconv.ParseInt(o["value_quantity"].(string),10,64)
                break;
            }
        }

        var total_quantity = 0;
        for _,o_interf := range d["options"].([]interface{}) {
            o := o_interf.(map[string]interface{})
            size_quantity,_ := strconv.ParseInt(o["value_quantity"].(string),10,64)
            o["value_quantity"] = strconv.FormatFloat((math.Min(float64(size_quantity),float64(color_quantity))),'f',-1,64)
            total_quantity += int(size_quantity)
        }
        d["quantity"] = strconv.FormatInt(int64(total_quantity),10);
    }

    json_bytes,_ := ffjson.Marshal(data)
    bodyStr := string(json_bytes)
    fmt.Fprintf(w, bodyStr)
}

func main() {
    runtime.GOMAXPROCS(16)
    http.HandleFunc("/", handleRequest) //设置访问的路由
    err := http.ListenAndServe(":18083", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
