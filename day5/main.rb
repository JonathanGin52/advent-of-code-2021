# frozen_string_literal: true

map = Hash.new { |h, k| h[k] = Hash.new(0) }
count = 0

File.readlines('input.txt').each do |line|
  start, dest = line.split(' -> ')

  x1, y1 = start.split(',').map(&:to_i)
  x2, y2 = dest.split(',').map(&:to_i)
  xs = [x1, x2]
  ys = [y1, y2]

  next unless x1 == x2 || y1 == y2

  (xs.min..xs.max).each do |x|
    (ys.min..ys.max).each do |y|
      map[x][y] += 1

      count += 1 if map[x][y] == 2
    end
  end
end

puts count
