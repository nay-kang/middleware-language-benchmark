package net.codeedu;

import java.io.IOException;
import java.io.OutputStream;
import java.net.URL;
import java.util.Iterator;

import javax.servlet.ServletException;
import javax.servlet.annotation.WebServlet;
import javax.servlet.http.HttpServlet;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import com.fasterxml.jackson.core.JsonFactory;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.node.ObjectNode;

/**
 * Servlet implementation class Index
 */
@WebServlet("/Index")
public class Index extends HttpServlet {
	private static final long serialVersionUID = 1L;

	/**
	 * Default constructor.
	 */
	public Index() {
		// TODO Auto-generated constructor stub
	}

	/**
	 * @see HttpServlet#doGet(HttpServletRequest request, HttpServletResponse
	 *      response)
	 */
	protected void doGet(HttpServletRequest request,
			HttpServletResponse response) throws ServletException, IOException {
		JsonFactory f = new JsonFactory();
		ObjectMapper m = new ObjectMapper();

		URL url = new URL("http://localhost:18082/product_list.json");
		JsonNode data = m.readTree(url.openStream());
		Iterator<JsonNode> list = data.path("list").elements();
		while (list.hasNext()) {
			JsonNode p = list.next();

			ObjectNode price = (ObjectNode) p.path("price");
			price.put("value", price.get("value").asDouble() * 1.25);

			int color_quantity = 0;
			Iterator<JsonNode> options = p.path("options").elements();
			while (options.hasNext()) {
				JsonNode o = options.next();
				if ("Color".equals(o.get("title").asText())) {
					color_quantity = o.get("value_quantity").asInt();
					break;
				}
			}
			options = p.path("options").elements();
			int total_quantity = 0;
			while (options.hasNext()) {
				JsonNode o = options.next();
				((ObjectNode) o).put("value_quantity", Math.min(
						o.get("value_quantity").asInt(), color_quantity));
				total_quantity += o.get("value_quantity").asInt();
			}

			((ObjectNode) p).put("quantity", total_quantity);
		}
		OutputStream ofs = response.getOutputStream();
		m.writeValue(ofs, data);

	}

}
