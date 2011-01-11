; Steve Phillips / elimisteve
; 2011.01.07

(eval
 `(+ 10
	 ~(* 3 2)))
; 16

(eval
 '(+ 10
	 (* 3 2)))
; 16

(eval
 (let [x '(+ 2 3)]
   `(~@x 1)))
; 6

(= `~'+ '+)
; true

;(defm )