SELECT id, customer_name, order_date, product, quantity, price
	FROM public.items
	WHERE quantity > 2 AND price > 50;