; Steve Phillips / elimisteve
; 2011.01.06

; Applies anonymous an function, which adds 2 to its input, to 5/2

(println ((fn [x] (inc (inc x))) 5/2))  ; 9/2

; Shortcut syntax, where '%' means the var passed in
(println (#(inc(inc %)) 5/2))  ; 9/2