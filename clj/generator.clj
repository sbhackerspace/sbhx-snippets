; Steve Phillips / elimisteve
; 2011.03.26

(defn generator []
  (def num 0)
  (fn []
    (def num (inc num))
    num))

(def g (generator))

(g)
(g)
(g)