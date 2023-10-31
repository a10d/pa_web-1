<?php

namespace Tests\Feature;

use App\Models\User;
use Illuminate\Foundation\Testing\RefreshDatabase;
use Illuminate\Support\Facades\Auth;
use Tests\TestCase;

class AuthTest extends TestCase
{
    use RefreshDatabase;

    public function test_a_user_can_log_in()
    {
        $credentials = [
            'username' => 'test',
            'password' => 'test',
        ];

        User::factory()->create($credentials);

        $this->assertFalse(Auth::check());

        $this->post('/api/auth/login', $credentials)->assertOk();

        $this->assertTrue(Auth::check());
    }

    public function test_a_user_can_list_other_users()
    {
        User::factory()->count(10)->create();

        $this
            ->actingAs(User::factory()->create())
            ->get('/api/auth/users')
            ->assertJsonCount(11, 'users');
    }
}
