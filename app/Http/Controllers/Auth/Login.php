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
        $authenticated = Auth::attempt($request->only('username', 'password'));

        if (!$authenticated) {
            abort(Response::HTTP_BAD_REQUEST);
        }
        $token = Auth::user()->createToken('authToken');

        return response()->json([
            'token' => $token->plainTextToken,
        ]);
    }

}
