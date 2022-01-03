# frozen_string_literal: true

class Board
  def initialize(board)
    @numbers = []
    board.split("\n").each { |line| @numbers << line.split.map(&:to_i) }
    @numbers.flatten!
    @won = false
  end

  def won?
    return true if @won

    slices = @numbers.each_slice(5)
    if slices.any? { |row| row.all?(&:negative?) }
      @won = true
      return true
    end
    if (0..4).map { |i| slices.map { |s| s[i] } }.any? { |column| column.all?(&:negative?) }
      @won = true
      return true
    end

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
    next if board.won?

    board.mark(number)
    puts board.score(number) if board.won?
  end
end
