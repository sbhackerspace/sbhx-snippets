#!/usr/bin/env ruby
# Steve Phillips / elimisteve
# 2011.06.21

# From http://ruby.about.com/od/gems/a/nokogiri.htm

require 'rubygems'
require 'nokogiri'
require 'open-uri'

doc = Nokogiri::HTML( open('http://www.google.com/search?q=nokogiri') )

doc.css('a.l').each do|l|
  puts l.content
end
