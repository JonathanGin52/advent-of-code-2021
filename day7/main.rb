# frozen_string_literal: true

positions = File.read('input.txt').split(',').map(&:to_i).sort
target = positions[positions.size / 2]
puts positions.sum { |loc| (loc - target).abs }
