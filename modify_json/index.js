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

            var data = JSON.parse(body);
            for(var key in data['list']){
                var d = data['list'][key];
                d['price']['value'] *=1.25;
                var color_quantity = 0;
                for(var o_key in d['options']){
                    var o = d['options'][o_key]
                    if(o['title'] == 'Color'){
                        color_quantity = o['value_quantity'];
                        break;
                    }
                }
                var total_quantity = 0;
                for(var o_key in d['options']){
                    var o = d['options'][o_key]
                    o['value_quantity'] = Math.min(o['value_quantity'],color_quantity);
                    total_quantity += o['value_quantity'];
                }
                d['quantity'] = total_quantity;
            }

            response.write(JSON.stringify(data));    
            response.end();
        });
    });
   response.writeHeader(200); 
}).listen(18083);
