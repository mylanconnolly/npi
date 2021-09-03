# frozen_string_literal: true

require_relative 'npi/version'

# This is the main NPI module. Currently, the only public method here is
# `valid?`, whch is used to check whether an NPI number is valid.
module NPI
  NPI_LENGTH = 10
  NUMERIC_REGEXP = /\A\d+\z/.freeze

  # This method is used to check whether an NPI number is valid. For our
  # purposes, "valid" is defined as having a matching check digit and matching
  # the defined NPI format.
  #
  # This does not check whether the NPI number has been assigned.
  def self.valid?(number)
    number = number.to_s
    return false if number.length < 10 || !number.match(NUMERIC_REGEXP)

    check_digit(number) == number[-1].to_i
  end

  def self.check_digit(number)
    check_val = check_value(number)
    check_digit = check_val % 10
    check_digit.zero? ? check_digit : 10 - check_digit
  end

  private_class_method :check_digit

  def self.check_value(number)
    number[0...9]
      .chars
      .map(&:to_i)
      .each_with_index
      .map { |v, i| i.even? ? v * 2 : v }
      .map { |v| v < 10 ? v : v.to_s.chars.map(&:to_i) }
      .flatten
      .push(24)
      .sum
  end

  private_class_method :check_value
end
