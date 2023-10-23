<?php

namespace App\Http\Controllers\TaskType;

use App\Foundation\Http\Controller;
use App\Models\TaskType;
use Illuminate\Http\JsonResponse;

class AllTaskTypes extends Controller
{
    public function __invoke(): JsonResponse
    {
        $this->authorize('viewAny', [TaskType::class]);

        return response()->json([
            'taskTypes' => TaskType::all(),
        ]);
    }

}
