<?php

namespace App\Http\Controllers\TaskType;

use App\Foundation\Http\Controller;
use App\Models\TaskType;

class AllTaskTypes extends Controller
{
    public function __invoke()
    {
        $this->authorize('viewAny', [TaskType::class]);

        return response()->json([
            'taskTypes' => TaskType::all(),
        ]);
    }

}
