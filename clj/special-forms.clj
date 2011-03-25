;;William Duncan 2/13/2011
;;Dicking around with special forms

;;Defines the namespace
(ns gtech)

;;Oh god halp! I'll have to try to decifer this...
(defn
 ^{:doc "mymax [xs+] gets the maximum value in xs using > ";;documentation, remember to fill this in!!! Not sure about the tilde though, maybe it'll come up later, for now LET THIS BE A NOTE
   :test (fn []
             (assert (= 42  (mymax 2 42 5 4))));testcase, clever
   :user/comment "this is the best fn ever!"};redundant comment key
  mymax
  ([x] x)
  ([x y] (if (> x y) x y));overloading
  ([x y & more]
   (reduce mymax (mymax x y) more)));sweet sexy recurssion
;;cool okay I get it, dear sweet raptorjesus is that beautiful

(doc println) ;;documentation for that bitch!!!

;So a let statement makes (a) local variable(s) for the duration of the combination, there is an implied do.
(let [x 3,
      y 6]
  (println (+ x y)))

;:as creates another variable with all of the values together
(let [[a b c d :as largenum]
      [1 2 3 4]]
  (println largenum)
  (println a "-" b c d))

(let [] 
;Do does a few actions in sequence
(do (println (* 3 5))
    (println "Hello"))

(defn hello
  "This function says hello!"
  [] (println "Hello there!"))