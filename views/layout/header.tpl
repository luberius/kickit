{{ define "header" }}
    <head>
        <meta content="text/html;charset=utf-8" http-equiv="Content-Type">
        <meta content="utf-8" http-equiv="encoding">
        <link rel="stylesheet" href="asset/main.css">
        <link rel="stylesheet" href="asset/tailwind.min.css">
    </head>
    <body>
    <div class="antialiased bg-gray-100 dark-mode:bg-gray-900">
        <div class="w-full text-gray-700 bg-blue-100 dark-mode:text-gray-200 dark-mode:bg-gray-800">
            <div class="flex flex-col max-w-screen-xl px-4 mx-auto md:items-center md:flex-row md:px-6 lg:px-8">
                <div class="flex flex-row items-center justify-between p-4">
                    <a href="#"
                       class="text-lg font-black tracking-widest text-gray-900 uppercase rounded-lg dark-mode:text-white">
                        Kick it!.
                    </a>
                </div>
                <div class="mx-2 w-1 h-5 border-gray-500 border-l-2 ml-10"></div>
                <span>Welcome</span>
            </div>
        </div>
    </div>
    <div class="flex flex-row max-w-screen-xl mx-auto pt-5">
        <!-- Menu -->
        <div class="w-64 h-screen bg-white">
            <nav>
                <div>
                    <button class="w-full flex justify-between items-center py-3 px-6 text-blue-600 cursor-default focus:outline-none">
                        <span class="flex items-center">
                            <svg class="h-5 w-5" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                                <path d="M19 11H5M19 11C20.1046 11 21 11.8954 21 13V19C21 20.1046 20.1046 21 19 21H5C3.89543 21 3 20.1046 3 19V13C3 11.8954 3.89543 11 5 11M19 11V9C19 7.89543 18.1046 7 17 7M5 11V9C5 7.89543 5.89543 7 7 7M7 7V5C7 3.89543 7.89543 3 9 3H15C16.1046 3 17 3.89543 17 5V7M7 7H17"
                                      stroke="currentColor" stroke-width="2" stroke-linecap="round"
                                      stroke-linejoin="round"></path>
                            </svg>
                            <span class="mx-4 font-medium">Project</span>
                        </span>
                    </button>

                    <div class="bg-white">
                        <a class="py-2 px-16 block text-sm text-gray-600 hover:text-gray-900" href="#">Create New
                            Project</a>
                        <a class="py-2 px-16 block text-sm text-gray-600 hover:text-gray-900" href="#">All Projects</a>
                    </div>
                </div>

                <div>
                    <button class="w-full flex justify-between items-center py-3 px-6 text-blue-600 cursor-default focus:outline-none">
                        <span class="flex items-center">
                           <svg xmlns='http://www.w3.org/2000/svg' width='24' height='24'
                                viewBox='0 0 512 512'><path
                                       d='M112,111V401c0,17.44,17,28.52,31,20.16l247.9-148.37c12.12-7.25,12.12-26.33,0-33.58L143,90.84C129,82.48,112,93.56,112,111Z'
                                       style='fill:none;stroke:#3182ce;stroke-miterlimit:10;stroke-width:32px'/></svg>

                            <span class="mx-4 font-medium">Active Project</span>
                        </span>
                    </button>

                    <div class="bg-white">
                        <a class="py-2 px-16 block text-sm font-semibold text-gray-800" href="#">(None)</a>
                        <a class="py-2 px-16 block text-sm text-gray-600 hover:text-gray-900" href="#">New Form</a>
                        <a class="py-2 px-16 block text-sm text-gray-600 hover:text-gray-900" href="#">All Form</a>
                        <a class="py-2 px-16 block text-sm text-gray-600 hover:text-gray-900" href="/test">Test</a>
                    </div>
                </div>
            </nav>
        </div>
{{ end }}