; Steve Phillips / elimisteve
; 2011.01.06

(defn fact1 [n]
  (loop [cnt n acc 1]
	(if (zero? cnt)
	  acc
	  (recur (dec cnt) (* acc cnt)))))

(println "(fact1 5) =>" (fact1 5))  ; 120


;; (defn fact2 [n]

;; (println "(fact2 5) => " (fact2 5))


(defn fact3 [n]
  (reduce *
		  (range 1 (inc n))))

(println "(fact3 5) =>" (fact3 5))