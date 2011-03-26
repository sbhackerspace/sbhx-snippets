; Steve Phillips / elimisteve
; 2011.03.26

(defn generator
  ([] (def num 0)
     (fn []
       (def num (inc num))
       num))
  ([x] (def num x)
     (fn []
       (def num (inc num))
       num)))

;; Usage
(def g1 (generator))
(g1)  ; 1
(g1)  ; 2
(g1)  ; 3

(def g2 (generator 10))
(g2)  ; 11
(g2)  ; 12
(g2)  ; 13
