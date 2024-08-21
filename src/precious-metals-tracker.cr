require "option_parser"
require "./../src/commands/add-holding.cr"

module Precious::Metals::Tracker
  VERSION = "0.0.0"

  OptionParser.parse do |parser|
    parser.banner = "Welcome To The Precious Metals Tracker!"

    parser.on "-v", "--version", "Show version" do
      puts "version: " + VERSION
      exit
    end

    parser.on "-h", "--help", "Show help" do
      puts parser

      exit
    end

    parser.on "-a", "--add", "Add new holding" do
      AddHolding.execute

      exit
    end

    parser.missing_option do |option_flag|
      STDERR.puts "ERROR: #{option_flag} is missing something."
      STDERR.puts ""
      STDERR.puts parser

      exit(1)
    end

    parser.invalid_option do |option_flag|
      STDERR.puts "ERROR: #{option_flag} is not a valid option."
      STDERR.puts parser

      exit(1)
    end
  end
end
