# frozen_string_literal: true

fish = File.read('input.txt').split(',').map(&:to_i)
days = 80

days.times do
  new_fish_count = 0
  fish.each_index do |i|
    fish[i] -= 1
    if fish[i] < 0
      fish[i] = 6
      new_fish_count += 1
    end
  end
  new_fish_count.times { fish << 8 }
end

puts fish.size
