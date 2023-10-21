<?php

namespace App\Http\Controllers\Auth;

use App\Foundation\Http\Controller;
use app\Http\Requests\Auth\Login as LoginRequest;
use Illuminate\Http\JsonResponse;
use Illuminate\Support\Facades\Auth;

class Login extends Controller
{
    public function __invoke(LoginRequest $request): JsonResponse
    {
        Auth::attempt($request->only('email', 'password'));

        $token = Auth::user()->createToken('authToken');

        return response()->json([
            'token' => $token->plainTextToken,
        ]);
    }

}
