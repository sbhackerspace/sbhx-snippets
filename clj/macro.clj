; Steve Phillips / elimisteve
; 2011.01.07

;; My first custom macro!!
(defn comp-nth
  "Compute f composed with itself n times of x"
  [n f x]
  (first (eval
		  `(map (comp ~@(repeat n f))
				(list ~x)))))

(println (comp-nth 5 #(+ 10 %) 0))


;; Better version; uses apply instead of no eval and backquotes
(defn compnth [n f x]
  (first (map
		  (apply comp (repeat n f))
		  (list x))))

(println (compnth 5 #(+ 10 %) 0))


;; 2011.01.20
;; No need for comp-nth/compnth/nth-comp -- use _iterate_ instead
(defn nth-comp [n f x]
  (nth (iterate f x) n))

(println (nth-comp 5 #(+ 10 %) 0))
