require "colorize"

module AddHolding
  extend self

  # Execute the command and handle the user inputs
  def execute
    self.request_information
  end

  def request_information
    puts "Please enter the holding information:"

    puts "Purchase Source > "
    purchase_source : String? = gets
    if purchase_source.nil? || purchase_source.blank?
      puts "no new holding provided".colorize(:red)

      exit
    end

    puts "#{user_inputs}"
  end
end
