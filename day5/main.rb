# frozen_string_literal: true

map = Hash.new { |h, k| h[k] = Hash.new(0) }
count = 0

File.readlines('input.txt').each do |line|
  start, dest = line.split(' -> ')

  x1, y1 = start.split(',').map(&:to_i)
  x2, y2 = dest.split(',').map(&:to_i)

  if x1 == x2 || y1 == y2
    # Horizontal and vertical lines
    xs = [x1, x2]
    ys = [y1, y2]
    (xs.min..xs.max).each do |x|
      (ys.min..ys.max).each do |y|
        map[x][y] += 1

        count += 1 if map[x][y] == 2
      end
    end
  elsif (y2 - y1).abs == (x2 - x1).abs
    # Diagonal lines
    x = x1
    y = y1
    x_incr = x1 < x2 ? 1 : -1
    y_incr = y1 < y2 ? 1 : -1

    loop do
      map[x][y] += 1
      count += 1 if map[x][y] == 2
      break if x == x2 && y == y2

      x += x_incr
      y += y_incr
    end
  end
end

puts count
