<?php
$data = json_decode(file_get_contents("http://localhost:18082/product_list.json"),true);
foreach($data['list'] as &$d){
	$d['price']['value'] *=1.25;
	$color_quantity = 0;
	foreach($d['options'] as $o){
		if($o['title'] == 'Color'){
			$color_quantity = $o['value_quantity'];
			break;
		}
	}
	$total_quantity = 0;
	foreach($d['options'] as &$o){
		$o['value_quantity'] = min($o['value_quantity'],$color_quantity);
		$total_quantity += $o['value_quantity'];
	}
	$d['quantity'] = $total_quantity;
}
echo json_encode($data);