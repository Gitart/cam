package main

import (
  "github.com/fatih/color"
)


var (

  // Base set
  ipmain  = "http://108.61.245.170/"

  // Set color
  Whites  = color.New(color.FgWhite).PrintlnFunc()
  FgMag   = color.New(color.FgHiYellow)
  Cyan    = color.New(color.Bold, color.FgCyan)
  FgRed   = color.New(color.Bold, color.FgRed).PrintlnFunc()
  FgGreen = color.New(color.Bold, color.FgGreen).PrintlnFunc()
  FgReds  = color.New(color.Bold, color.FgRed)
  
  hlp     =`
            <html>
            <head> 
                  <link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" >
                  <script src="https://code.jquery.com/jquery-1.12.4.js"></script>
                  <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/js/bootstrap.min.js" ></script>
            </head>
            
            <body>      
            <div class="container">
                 <h1>Help usage api cam</h1>
                 <hr>
                 <p>
                    Использование камеры для получения снимков из Флориды.
                    Тестовое задание.
                 </p>

                 <h4></h4>
                 <hr>
                 <a href="/api/cam">Web cam</a>

            </div>
            </body>
            </html>

                 `
)
