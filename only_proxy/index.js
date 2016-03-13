var http=require("http");

http.createServer(function(request,response){
    http.get({
        host:"localhost",
        port:"18082",
        path:"/product_list.json"
    },function(resp){
        var body='';
        resp.on("data",function(d){
            body+=d;  
        });
        resp.on("end",function(){
            response.write(body);    
            response.end();
        });
    });
   response.writeHeader(200); 
}).listen(18083);
