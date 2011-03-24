; Steve Phillips / elimisteve
; 2011.01.26

(defn dot-product [v1 v2]
  (reduce + (map * v1 v2)))

;; I love Clojure...