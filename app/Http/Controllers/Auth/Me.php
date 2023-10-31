<?php

namespace App\Http\Controllers\Auth;

use App\Foundation\Http\Controller;
use Illuminate\Support\Facades\Auth;

class Me extends Controller
{
    public function __invoke()
    {
        $this->middleware('auth');

        $user = Auth::user();

        return [
            'id' => Auth::id(),
            'name' => $user->name,
            'username' => $user->username,
        ];
    }

}
