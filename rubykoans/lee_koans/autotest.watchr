def run_rake_test
  system('rake')
end

watch( '(koans)/.*\.rb' ) { run_rake_test }
