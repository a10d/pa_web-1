<?php

namespace Tests\Feature;

use App\Models\Task;
use App\Models\TaskType;
use App\Models\User;
use Illuminate\Foundation\Testing\RefreshDatabase;
use Illuminate\Http\Response;
use Illuminate\Support\Arr;
use Tests\TestCase;

class TaskTest extends TestCase
{
    use RefreshDatabase;

    public function test_a_guest_cannot_list_any_tasks()
    {
        $this->get('/api/tasks/all')->assertStatus(Response::HTTP_FORBIDDEN);
    }

    public function test_any_user_can_list_all_tasks()
    {
        $this->actingAs(User::factory()->create());

        Task::factory()->count(5)->create();

        $this->get('/api/tasks/all')
            ->assertOk()
            ->assertJsonCount(5, 'tasks');
    }

    public function test_a_user_can_create_a_task()
    {
        $data = [
            'type' => TaskType::factory()->create()->id,
            'title' => 'Test Task',
            'description' => 'Test Task Description',
            'dueDate' => now()->addDay()->format('Y-m-d H:i:s'),
            'assignees' => [User::factory()->create()->id]
        ];

        $this
            ->actingAs(User::factory()->create())
            ->post('/api/tasks/create', $data)
            ->assertOk();

        $this->assertDatabaseHas('tasks', [
            'task_type_id' => $data['type'],
            ...Arr::except($data, ['type', 'assignees'])
        ]);
    }

    public function test_a_user_can_update_a_task()
    {
        $task = Task::factory()->create([
            'title' => 'Task before update',
        ]);

        $this->assertEquals('Task before update', $task->title);

        $this
            ->actingAs(User::factory()->create())
            ->patch('/api/tasks/update/' . $task->id, [
                ...Arr::only($task->toArray(), ['type', 'title', 'description', 'dueDate', 'completed']),
                'type' => $task->taskType->id,
                'title' => 'Task after update',
            ])
            ->assertOk();

        $task->refresh();

        $this->assertEquals('Task after update', $task->title);
    }

    public function test_a_user_can_delete_an_existing_task()
    {
        $task = Task::factory()->create();

        $this
            ->actingAs(User::factory()->create())
            ->delete('/api/tasks/delete/' . $task->id)
            ->assertOk();

        $this->assertDatabaseMissing('tasks', [
            'id' => $task->id,
        ]);
    }

    public function test_the_assignees_of_a_task_can_be_changed()
    {

        $task = Task::factory()->create();

        $this->assertEmpty($task->assignees);

        $user1 = User::factory()->create();
        $user2 = User::factory()->create();

        $this
            ->actingAs(User::factory()->create())
            ->patch('/api/tasks/update/' . $task->id, [
                ...Arr::only($task->toArray(), ['title', 'description', 'dueDate', 'completed']),
                'assignees' => [$user1->id, $user2->id],
            ])
            ->assertOk();

        $task->refresh();

        $this->assertCount(2, $task->assignees);
        $this->assertContains($user1->id, $task->assignees);
        $this->assertContains($user2->id, $task->assignees);
    }

    public function test_an_assignee_of_a_task_can_be_removed()
    {
        $task = Task::factory()->create();

        $user1 = User::factory()->create();
        $user2 = User::factory()->create();

        $task->assignees = [$user1->id, $user2->id];

        $this->assertCount(2, $task->assignees);

        $this
            ->actingAs(User::factory()->create())
            ->patch('/api/tasks/update/' . $task->id, [
                ...Arr::only($task->toArray(), ['title', 'description', 'dueDate', 'completed']),
                'assignees' => [$user1->id],
            ])
            ->assertOk();

        $task->refresh();

        $this->assertCount(1, $task->assignees);
        $this->assertContains($user1->id, $task->assignees);
        $this->assertNotContains($user2->id, $task->assignees);
    }
}
