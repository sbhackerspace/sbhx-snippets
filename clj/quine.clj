; Steve Phillips / elimisteve
; 2011.01.11

;;seriously wtf is this?
(defn quine [x]
  (fn [x](list x
	       (list (quote quote)
		     x)))
  (quote (fn [x] (list x
		       (list (quote quote)
			     x)))))


