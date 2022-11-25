#{Define .a.1 = file1}
#{Define .a.2 = file2}
#{Define .a.3 = file3}

#{Loop #{.a.*Length} index=a.index}
#{  StartFile  }
#{  Define file.name=#{GetKey .a #{a.index}}  }
#{  Define .c=#{file.name}  }
#{  Insert header.i  }
hello world!
#{  EndFile  }
#{EndLoop}