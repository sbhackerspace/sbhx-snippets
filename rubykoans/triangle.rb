# Triangle Project Code.

# Triangle analyzes the lengths of the sides of a triangle
# (represented by a, b and c) and returns the type of triangle.
#
# It returns:
#   :equilateral  if all sides are equal
#   :isosceles    if exactly 2 sides are equal
#   :scalene      if no sides are equal
#
# The tests for this method can be found in
#   about_triangle_project.rb
# and
#   about_triangle_project_2.rb
#
def triangle(a, b, c)
  # Side lengths must be positive
  if !(a > 0 and b > 0 and c > 0)
    raise TriangleError
  end

  # Triangle inequality must hold
  if !(a + b > c and a + c > b and b + c > a)
    raise TriangleError
  end

  if a == b and b == c
    :equilateral
  elsif
    (a == b and b != c) or (a == c and c != b) or (b == c and c != a)
    :isosceles
  else
    :scalene
  end
end

# Error class used in part 2.  No need to change this code.
class TriangleError < StandardError
end
