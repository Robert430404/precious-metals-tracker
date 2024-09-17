require "colorize"

module Commands::AddHolding
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
      puts "please enter a valid purchase source".colorize(:red)

      exit
    end

    puts "Purchase Price > "
    purchase_price : String? = gets
    if purchase_source.nil? || purchase_source.blank?
      puts "please enter a valid purchase price".colorize(:red)

      exit
    end

    puts "Spot Price At Time Of Purchase > "
    purchase_spot_price : String? = gets
    if purchase_spot_price.nil? || purchase_spot_price.blank?
      puts "please enter a valid purchase spot price".colorize(:red)

      exit
    end

    puts "#{purchase_source}"
    puts "#{purchase_price}"
    puts "#{purchase_spot_price}"
  end
end
