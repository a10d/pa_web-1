<?php

namespace App\Http\Controllers\Task;

use App\Foundation\Http\Controller;
use App\Models\Task;
use Illuminate\Http\JsonResponse;

class AllTasks extends Controller
{
    public function __invoke(): JsonResponse
    {
        $this->authorize('viewAny', [Task::class]);

        return response()->json([
            'tasks' => Task::all(),
        ]);
    }

}
