def showNbyN(n):
    for i in range(n):
        print [(x+i)%n+1 for x in range(n)]
        

showNbyN(5)
