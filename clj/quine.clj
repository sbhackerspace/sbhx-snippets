; Steve Phillips / elimisteve
; 2011.01.11

((fn [x]
   (list x
		 (list (quote quote)
			   x)))
 (quote
  (fn [x]
	(list x
		  (list (quote quote)
				x)))))
