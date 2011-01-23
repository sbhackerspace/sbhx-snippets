; Steve Phillips / elimisteve
; 2011.01.22

(interleave (range 10) (range 10 20))
;; (0 10 1 11 2 12 3 13 4 14 5 15 6 16 7 17 8 18 9 19)

;; Turn it into a hash map with assoc
(def mymap (apply assoc {} (interleave (range 10) (range 10 20))))
;; {0 10, 1 11, 2 12, 3 13, 4 14, 5 15, 6 16, 7 17, 8 18, 9 19}

(mymap 2)
;; 12
