<?php

namespace App\Http\Controllers\Auth;

use App\Foundation\Http\Controller;
use App\Http\Requests\Auth\Login as LoginRequest;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Response;
use Illuminate\Support\Facades\Auth;

class Login extends Controller
{
    public function __invoke(LoginRequest $request): JsonResponse
    {
        if (Auth::attempt($request->only('username', 'password'), true)) {
            return response()->json([
                'message' => 'Successfully logged in',
            ]);
        }

        return response()->json([
            'message' => 'Invalid credentials',
        ])->setStatusCode(Response::HTTP_UNAUTHORIZED);
    }

}
