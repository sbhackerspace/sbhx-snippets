; Steve Phillips / elimisteve
; 2011.01.06

; Inspired by
;  (map char (concat (range 65 91) (range 97 123))) 
; at http://stackoverflow.com/questions/2578233/
; how-do-i-get-the-set-of-all-letters-in-java-clojure

; Uppercase
(println
 (map char (range (int \A) (inc (int \Z)))))

; Lowercase
(println
 (take 26 (map char (iterate inc (int \a)))))

; Uppercase and lowercase
(println
 (map char (concat
			(range (int \A) (inc (int \Z)))
			(take 26 (map char (iterate inc (int \a)))))))