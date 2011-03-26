;;William Duncan 3/26/11

(defn counter [x]
  (def x (inc x))
  x)
; Lexical scoping (right?)
(counter 5)  ; 5


(defn counter-binding [x]
  (binding [x (inc x)])
  x)  ; x comes after 'binding'

(counter-binding 5)  ; 5


(defn counter-binding2 [x]
  (binding [x (inc x)]
    x))  ; x inside 'binding'

(counter-binding2 5)  ; 5


(defn counter-let [x]
  (let [x (inc x)]
  x))

(counter-let 5)  ; 6


;; This is gtech's original function
(defn counter-fn [x]
  (fn []
    (def x (inc x))
    x))

((counter-fn 5))  ; 5
