<?php

namespace App\Http\Controllers\Auth;

use App\Foundation\Http\Controller;
use App\Models\User;
use Illuminate\Http\JsonResponse;

class AllUsers extends Controller
{
    public function __invoke(): JsonResponse
    {
        $this->middleware('auth');

        return response()->json([
            'users' => User::all()->map(fn(User $user) => [
                'id' => $user->id,
                'username' => $user->username,
            ]),
        ]);
    }
}
