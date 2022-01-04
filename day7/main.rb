# frozen_string_literal: true

positions = File.read('input.txt').split(',').map(&:to_i).sort

min_sum = Float::INFINITY
positions.min.upto(positions.max).each do |target|
  sum = positions.sum do |loc|
    delta = (loc - target).abs
    delta * (delta + 1) / 2
  end
  min_sum = [sum, min_sum].min
end

puts min_sum
