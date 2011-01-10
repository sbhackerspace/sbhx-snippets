; Steve Phillips / elimisteve
; 2011.01.09

; Extract phone number
(apply str (re-seq #"\d{3,4}" "(805) 246-1234"))  ; "8052461234"
(apply str (re-seq #"\d{3,4}" "805-123-5678"))  ; "8051235678"

; Extract first name from email address
(re-find #"[a-zA-Z]+" "bob.k.jones@yahoo.com")  ; "bob"
