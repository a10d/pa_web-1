<!DOCTYPE html>
<html lang="{{ str_replace('_', '-', app()->getLocale()) }}">
<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">

    <title>Aufgabenverwaltung</title>
    @vite('resources/css/app.css')
    @vite('resources/js/index.ts')
</head>
<body class="antialiased">
<div id="app">
    <div class="flex fixed inset-0">
        <div class="m-auto grid grid-cols-2 gap-2 h-4 w-4 animate-spin">
            @for($i=0;$i<4;$i++)
                <div class="w-1 h-1 bg-slate-300 rounded-full"></div>
            @endfor
        </div>
    </div>
</div>
</body>
</html>
