<?php

namespace Tests\Feature;

use App\Models\TaskType;
use App\Models\User;
use Illuminate\Foundation\Testing\RefreshDatabase;
use Illuminate\Support\Arr;
use Tests\TestCase;

class TaskTypeTest extends TestCase
{
    use RefreshDatabase;

    public function test_an_authenticated_user_can_list_all_task_types()
    {
        $this->actingAs(User::factory()->create());

        TaskType::factory()->count(5)->create();

        $this->get('/api/taskTypes/all')->assertStatus(200)->assertJsonCount(5, 'taskTypes');
    }

    public function test_a_task_type_can_be_created()
    {
        $user = User::factory()->create(['username' => 'admin']);

        $this->actingAs($user);

        $data = [
            'name' => 'Test Task Type',
            'description' => 'Test Task Type Description',
            'color' => '#000000',
            'reminderTime' => 2,
        ];

        $this->post('/api/taskTypes/create', $data)->assertStatus(200);

        $this->assertDatabaseHas('task_types', $data);
    }

    public function test_a_default_user_cannot_create_a_task()
    {
        $user = User::factory()->create(['username' => 'user']);

        $this->actingAs($user);

        $data = [
            'name' => 'Test Task Type',
            'description' => 'Test Task Type Description',
            'color' => '#000000',
            'reminderTime' => 2,
        ];

        $this->post('/api/taskTypes/create', $data)->assertStatus(403);

        $this->assertDatabaseMissing('task_types', $data);
    }


    public function test_an_anonymous_guest_cannot_create_a_task_type()
    {
        $data = [
            'name' => 'Test Task Type',
            'description' => 'Test Task Type Description',
            'color' => '#000000',
            'reminderTime' => 2,
        ];

        $this->post('/api/taskTypes/create', $data)->assertStatus(403);

        $this->assertDatabaseMissing('task_types', $data);
    }

    public function test_an_admin_can_update_task_types()
    {
        $taskType = TaskType::factory()->create([
            'name' => 'Before Change'
        ]);

        $user = User::factory()->create(['username' => 'admin']);

        $this->actingAs($user);

        $data = [
            ...Arr::only($taskType->toArray(), ['description', 'color', 'reminderTime']),
            'name' => 'After Change'
        ];

        $this->patch('/api/taskTypes/update/' . $taskType->id, $data)->assertStatus(200);

        $this->assertDatabaseHas('task_types', $data);
    }

    public function test_an_admin_can_delete_a_task_type()
    {
        $user = User::factory()->create(['username' => 'admin']);

        $this->actingAs($user);

        $taskType = TaskType::factory()->create();

        $this->delete('/api/taskTypes/delete/' . $taskType->id)->assertStatus(200);

        $this->assertDatabaseMissing('task_types', $taskType->toArray());

    }
}
