# frozen_string_literal: true

puts (File.readlines('input.txt').sum do |line|
  line.split(' | ')[1].split(' ').count { |o| [2, 4, 3, 7].include?(o.length) }
end)
