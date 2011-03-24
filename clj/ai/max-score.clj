; Steve Phillips / elimisteve
; 2011.01.26

;; max-key can be used to calculate the optimal choice for each player

;; Remember that vectors are functions of their indices
([1 2 3] 0)
; 1

;; Which vector has the largest first element (subscript 0)?
(max-key #(% 0) [1 5] [6 0] [3 3])
; [6 0]

;; Second element (subscript 1)?
(max-key #(% 1) [1 5] [6 0] [3 3])
; [1 5]

;; The above works regardless of the length of each vector.  For
;; Google's AI competition, our tree elements will probably always be
;; length 2, in which case we can do the following --

(min-key first [1 7] [3 5] [4 4])
; [1 7]

(max-key first [1 7] [3 5] [4 4])
; [4 4]

(group-by #(apply max %) [[1 2] [9 8]])
; {2 [[1 2]], 9 [[9 8]]}
