require "colorize"

module Commands::AddHolding
  extend self

  # Execute the command and handle the user inputs
  def execute
    self.request_information
  end

  def request_field(field_name : String) : String
    puts "#{field_name} > "
    result : String? = gets
    if result.nil? || result.blank?
      puts "please enter a valid #{field_name}".colorize(:red)
      exit
    end

    return result
  end

  def request_information
    puts "Please enter the holding information:"

    purchase_source : String = self.request_field("Purchase Source")
    purchase_price : String = self.request_field("Purchase Price")
    purchase_spot_price : String = self.request_field("Spot Price At Time Of Purchase")
    total_oz : String = self.request_field("Total Weight In Troy Ounces")

    puts "#{purchase_source} - #{purchase_price} - #{purchase_spot_price} - #{total_oz}"
  end
end
