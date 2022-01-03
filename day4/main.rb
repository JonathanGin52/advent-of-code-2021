# frozen_string_literal: true

class Board
  def initialize(board)
    @numbers = []
    board.split("\n").each { |line| @numbers << line.split.map(&:to_i) }
    @numbers.flatten!
  end

  def won?
    slices = @numbers.each_slice(5)
    return true if slices.any? { |row| row.all?(&:negative?) }
    return true if (0...5).map { |i| slices.map { |s| s[i] } }.any? { |column| column.all?(&:negative?) }

    false
  end

  def mark(drawn_number)
    @numbers.flatten.each.each_with_index do |num, index|
      @numbers[index] = -num if num == drawn_number
    end
  end

  def score(just_called)
    @numbers.find_all(&:positive?).sum * just_called
  end

  def to_s
    @numbers.each_slice(5).reduce('') do |sum, slice|
      sum + "#{slice.map(&:abs).join(' ')}\n"
    end
  end
end

data = File.read('input.txt').split("\n\n")

numbers = data.shift.split(',').map(&:to_i)
boards = data.map { |b| Board.new(b) }

numbers.each do |number|
  boards.each do |board|
    board.mark(number)

    if board.won?
      puts board.score(number)
      exit
    end
  end
end
