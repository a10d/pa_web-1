<?php

namespace Tests\Feature;

use App\Models\User;
use Illuminate\Foundation\Testing\RefreshDatabase;
use Tests\TestCase;

class TaskTypeTest extends TestCase
{
    use RefreshDatabase;

    public function test_a_task_type_can_be_created()
    {
        $user = User::factory()->create();

        $this->actingAs($user);



    }
}
