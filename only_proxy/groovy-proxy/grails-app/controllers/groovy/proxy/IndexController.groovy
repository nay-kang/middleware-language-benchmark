package groovy.proxy

import groovyx.net.http.HTTPBuilder
import groovyx.net.http.ContentType

class IndexController {

    def index() {
        def http = new HTTPBuilder("http://localhost:18082")

        http.get(path:"/product_list.json",contentType:ContentType.TEXT){resp,reader->
            render reader.text
        }    
    }
}
