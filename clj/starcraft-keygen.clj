;;William Duncan 3/24/11
(ns starcraft-keygen)

(defn vector-of-random-digits
  "Return a vector of random digits. Params: max is the largest possible int from 0 to (n-1), length of vector."
  [max length]
  (let [l (dec length)];;Is there a better way to do this? I want to avoid any possibility of infinite loops so I want to use >, but I only want to decriment the upper bound once.
   (loop [i 0, v []]
	 (if (> i l)
	   v
	   (recur (inc i)
		  (assoc v i (rand-int max)))))))
  

(defn gen-key
  "Generates and returns a Starcraft key"
  []
  (def first-12-digits
    (vector-of-random-digits 10 12))
  
  
  
