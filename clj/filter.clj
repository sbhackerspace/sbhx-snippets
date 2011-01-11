; Steve Phillips / elimisteve
; 2011.01.06

; Using fn
(println (filter
		  (fn [x] (> x 3))
		  (range -10 5)))
; (4)

; Short-cut
(println (filter
		  #(> % 3)
		  (range -10 5)))
; (4)
