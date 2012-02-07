#!/usr/bin/env ruby
# Steve Phillips / elimisteve
# 2011.12.15

class FizzBuzz
  def self.solve
    for num in 1..100
      if num % 3 != 0 && num % 5 != 0
        puts num
        next
      end
      if num % 3 == 0
        print 'Fizz'
      end
      if num % 5 == 0
        print 'Buzz'
      end
      puts ''
    end
  end
end

FizzBuzz.solve
