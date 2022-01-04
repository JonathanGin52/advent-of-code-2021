# frozen_string_literal: true

data = File.read('input.txt').split(',').map(&:to_i)
days = 256

fish = Array.new(9, 0)
data.each { |timer| fish[timer] += 1 }

days.times do
  new_fish = fish[0]
  (0..7).each { |i| fish[i] = fish[i + 1] }
  fish[6] += new_fish
  fish[8] = new_fish
end

puts fish.sum
